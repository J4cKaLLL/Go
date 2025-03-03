package cache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// mockResult has the float64 and err to return
type mockResult struct {
	price float64
	err   error
}

type mockPriceService struct {
	numCalls    int
	mockResults map[string]mockResult // what price and err to return for a particular itemCode
	callDelay   time.Duration         // how long to sleep on each call so that we can simulate calls to be expensive
}

func (m *mockPriceService) GetPriceFor(itemCode string) (float64, error) {

	m.numCalls++            // increase the number of calls
	time.Sleep(m.callDelay) // sleep to simulate expensive call

	result, ok := m.mockResults[itemCode]
	if !ok {
		panic(fmt.Errorf("bug in the tests, we didn't have a mock result for [%v]", itemCode))
	}
	return result.price, result.err
}

func (m *mockPriceService) getNumCalls() int {
	return m.numCalls
}

func getPriceWithNoErr(t *testing.T, cache *TransparentCache, itemCode string) float64 {
	price, err := cache.GetPriceFor(itemCode)
	if err != nil {
		t.Error("error getting price for", itemCode)
	}
	return price
}

func getPricesWithNoErr(t *testing.T, cache *TransparentCache, itemCodes ...string) []float64 {
	prices, err := cache.GetPricesFor(itemCodes...)
	if err != nil {
		t.Error("error getting prices for", itemCodes)
	}
	return prices
}

func assertInt(t *testing.T, expected int, actual int, msg string) {
	if expected != actual {
		t.Error(msg, fmt.Sprintf("expected : %v, got : %v", expected, actual))
	}
}

func assertFloat(t *testing.T, expected float64, actual float64, msg string) {
	if expected != actual {
		t.Error(msg, fmt.Sprintf("expected : %v, got : %v", expected, actual))
	}
}

func assertFloats(t *testing.T, expected []float64, actual []float64, msg string) {
	if len(expected) != len(actual) {
		t.Error(msg, fmt.Sprintf("expected : %v, got : %v", expected, actual))
		return
	}
	for i, expectedValue := range expected {
		if expectedValue != actual[i] {
			t.Error(msg, fmt.Sprintf("result position [%v], expected : %v, got : %v", i, expected, actual))
			return
		}
	}
}

// Check that we are caching results (we should not call the external service for all calls)
func TestGetPriceFor_CachesResults(t *testing.T) {
	mockService := &mockPriceService{
		mockResults: map[string]mockResult{
			"p1": {price: 5, err: nil},
		},
	}
	cache := NewTransparentCache(mockService, time.Minute)
	assertFloat(t, 5, getPriceWithNoErr(t, cache, "p1"), "wrong price returned")
	assertFloat(t, 5, getPriceWithNoErr(t, cache, "p1"), "wrong price returned")
	assertFloat(t, 5, getPriceWithNoErr(t, cache, "p1"), "wrong price returned")
	assertInt(t, 1, mockService.getNumCalls(), "wrong number of service calls")
}

// Check that cache returns an error if external service returns an error
func TestGetPriceFor_ReturnsErrorOnServiceError(t *testing.T) {
	mockService := &mockPriceService{
		mockResults: map[string]mockResult{
			"p1": {price: 0, err: fmt.Errorf("some error")},
		},
	}
	cache := NewTransparentCache(mockService, time.Minute)
	_, err := cache.GetPriceFor("p1")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

// Check that cache can return more than one price at once, caching appropriately
func TestGetPricesFor_GetsSeveralPricesAtOnceAndCachesThem(t *testing.T) {
	mockService := &mockPriceService{
		mockResults: map[string]mockResult{
			"p1": {price: 5, err: nil},
			"p2": {price: 7, err: nil},
		},
	}
	cache := NewTransparentCache(mockService, time.Minute)
	assertFloat(t, 5, getPriceWithNoErr(t, cache, "p1"), "wrong price returned")
	assertFloats(t, []float64{5, 7}, getPricesWithNoErr(t, cache, "p1", "p2"), "wrong price returned")
	assertFloats(t, []float64{5, 7}, getPricesWithNoErr(t, cache, "p1", "p2"), "wrong price returned")
	assertInt(t, 2, mockService.getNumCalls(), "wrong number of service calls")
}

// Check that we are expiring results when they exceed the max age
func TestGetPriceFor_DoesNotReturnOldResults(t *testing.T) {
	mockService := &mockPriceService{
		mockResults: map[string]mockResult{
			"p1": {price: 5, err: nil},
			"p2": {price: 7, err: nil},
		},
	}
	maxAge := time.Millisecond * 200
	maxAge70Pct := time.Millisecond * 140
	cache := NewTransparentCache(mockService, maxAge)
	// get price for "p1" twice (one external service call)
	assertFloat(t, 5, getPriceWithNoErr(t, cache, "p1"), "wrong price returned")
	assertFloat(t, 5, getPriceWithNoErr(t, cache, "p1"), "wrong price returned")
	assertInt(t, 1, mockService.getNumCalls(), "wrong number of service calls")
	// sleep 0.7 the maxAge
	time.Sleep(maxAge70Pct)
	// get price for "p1" and "p2", only "p2" should be retrieved from the external service (one more external call)
	assertFloat(t, 5, getPriceWithNoErr(t, cache, "p1"), "wrong price returned")
	assertFloat(t, 5, getPriceWithNoErr(t, cache, "p1"), "wrong price returned")
	assertFloat(t, 7, getPriceWithNoErr(t, cache, "p2"), "wrong price returned")
	assertFloat(t, 7, getPriceWithNoErr(t, cache, "p2"), "wrong price returned")
	assertInt(t, 2, mockService.getNumCalls(), "wrong number of service calls")
	// sleep 0.7 the maxAge
	time.Sleep(maxAge70Pct)
	// get price for "p1" and "p2", only "p1" should be retrieved from the cache ("p2" is still valid)
	assertFloat(t, 5, getPriceWithNoErr(t, cache, "p1"), "wrong price returned")
	assertFloat(t, 5, getPriceWithNoErr(t, cache, "p1"), "wrong price returned")
	assertFloat(t, 7, getPriceWithNoErr(t, cache, "p2"), "wrong price returned")
	assertInt(t, 3, mockService.getNumCalls(), "wrong number of service calls")
}

// Check that cache parallelize service calls when getting several values at once
func TestGetPricesFor_ParallelizeCalls(t *testing.T) {
	mockService := &mockPriceService{
		callDelay: time.Second, // each call to external service takes one full second
		mockResults: map[string]mockResult{
			"p1": {price: 5, err: nil},
			"p2": {price: 7, err: nil},
		},
	}
	cache := NewTransparentCache(mockService, time.Minute)
	start := time.Now()
	assertFloats(t, []float64{5, 7}, getPricesWithNoErr(t, cache, "p1", "p2"), "wrong price returned")
	elapsedTime := time.Since(start)
	if elapsedTime > (1200 * time.Millisecond) {
		t.Error("calls took too long, expected them to take a bit over one second, perhaps calls were not parallelized?")
	}
}

// Check that cache returns an error when getting several values at once if one returns an error
// Results should be returned regardless, with a "0" for the one that gave an error
func TestGetPricesFor_ParallelizeCalls_Error(t *testing.T) {
	mockService := &mockPriceService{
		callDelay: time.Second, // each call to external service takes one full second
		mockResults: map[string]mockResult{
			"p1": {price: 5, err: nil},
			"p2": {price: 7, err: fmt.Errorf("some error")},
			"p3": {price: 9, err: nil},
		},
	}
	cache := NewTransparentCache(mockService, time.Minute)
	start := time.Now()
	results, err := cache.GetPricesFor("p1", "p2", "p3")
	assertFloats(t, []float64{5.0, 0.0, 9.0}, results, "wrong price returned")
	if err == nil {
		t.Error("expected call to GetPricesFor to return an error but nil was returned")
	}
	elapsedTime := time.Since(start)
	if elapsedTime > (1200 * time.Millisecond) {
		t.Error("calls took too long, expected them to take a bit over one second, perhaps calls were not parallelized?")
	}
}

// Check that if the same product code is requested from the cache concurrently
//
//	then the external service is not called twice at the same time
func TestGetPrices_DoesNotCallExternalServiceConcurrentlyForSameItem(t *testing.T) {
	mockService := &mockPriceService{
		callDelay: time.Second, // each call to external service takes one full second
		mockResults: map[string]mockResult{
			"p1": {price: 5, err: nil},
			"p2": {price: 7, err: nil},
		},
	}
	cache := NewTransparentCache(mockService, time.Minute)
	start := time.Now()
	// make three concurrent calls to the cache, two of them for the same item
	price1, price2, price3, err1, err2, err3 := 0.0, 0.0, 0.0, error(nil), error(nil), error(nil)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		price1, err1 = cache.GetPriceFor("p1")
		wg.Done()
	}()
	go func() {
		price2, err2 = cache.GetPriceFor("p1")
		wg.Done()
	}()
	go func() {
		price3, err3 = cache.GetPriceFor("p2")
		wg.Done()
	}()
	wg.Wait()
	assertFloats(t, []float64{5.0, 5.0, 7.0}, []float64{price1, price2, price3}, "wrong price returned")
	if err1 != nil || err2 != nil || err3 != nil {
		t.Error("did not expect the calls to return an error", err1, err2, err3)
	}
	elapsedTime := time.Since(start)
	if elapsedTime > (1200 * time.Millisecond) {
		t.Error("calls took too long, expected them to take a bit over one second, perhaps calls were not parallelized?")
	}
	if mockService.numCalls != 2 {
		t.Error("expecting only 2 calls to the service")
	}
}
