package service

import (
	"context"
	test "delivery/pkg/api/test/api"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	test.UnimplementedOrderServiceServer

	mu     sync.Mutex
	orders map[string]*test.Order
}

func New() *Service {
	return &Service{}
}

func (s *Service) CreateOrder(ctx context.Context, req *test.CreateOrderRequest) (*test.CreateOrderResponse, error) {

	orderId := uuid.New().String()
	order := test.Order{
		Id:       orderId,
		Item:     req.GetItem(),
		Quantity: req.GetQuantity(),
	}

	s.mu.Lock()
	s.orders[orderId] = &order
	s.mu.Unlock()

	return &test.CreateOrderResponse{Id: orderId}, nil
}

func (s *Service) GetOrder(ctx context.Context, req *test.GetOrderRequest) (*test.GetOrderResponse, error) {
	reqId := req.GetId()

	s.mu.Lock()
	respOrder, ok := s.orders[reqId]
	s.mu.Unlock()

	if !ok {
		return nil, status.Errorf(codes.NotFound, "No order with such ID: %s", reqId)
	}

	return &test.GetOrderResponse{Order: respOrder}, nil
}

func (s *Service) UpdateOrder(ctx context.Context, req *test.UpdateOrderRequest) (*test.UpdateOrderResponse, error) {
	reqId := req.GetId()

	// newOrder := test.Order{
	// 	Id:       req.Id,
	// 	Item:     req.Item,
	// 	Quantity: req.Quantity,
	// }

	s.mu.Lock()
	defer s.mu.Unlock()
	order, ok := s.orders[reqId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "No order with such ID: %s", reqId)
	}

	order.Item = req.GetItem()
	order.Quantity = req.GetQuantity()

	return &test.UpdateOrderResponse{Order: order}, nil
}

func (s *Service) DeleteOrder(ctx context.Context, req *test.DeleteOrderRequest) (*test.DeleteOrderResponse, error) {
	reqId := req.GetId()

	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.orders[reqId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "No order with such ID: %s", reqId)
	}
	delete(s.orders, reqId)

	return &test.DeleteOrderResponse{Success: true}, nil
}

func (s *Service) ListOrders(ctx context.Context, req *test.ListOrdersRequest) (*test.ListOrdersResponse, error) {
	respOrders := make([]*test.Order, 0)

	s.mu.Lock()
	defer s.mu.Unlock()

	for _, val := range s.orders {
		respOrders = append(respOrders, val)
	}

	return &test.ListOrdersResponse{Orders: respOrders}, nil
}
