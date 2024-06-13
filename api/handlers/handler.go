package handlers

import (
	"github.com/Project_Restaurant/api-gateway/genproto/menu"
	"github.com/Project_Restaurant/api-gateway/genproto/payment"
	"github.com/Project_Restaurant/api-gateway/genproto/reservation"
	"github.com/Project_Restaurant/api-gateway/genproto/restaurant"

	"google.golang.org/grpc"
)

type HandlerStruct struct {
	Payment     payment.PaymentServiceClient
	Reservation reservation.ReservationServiceClient
	Restaurant  restaurant.RestaurantServiceClient
	Menu        menu.MenuServiceClient
}

func NewHandler(conPay, conReser *grpc.ClientConn) *HandlerStruct {
	return &HandlerStruct{
		Payment:     payment.NewPaymentServiceClient(conPay),
		Reservation: reservation.NewReservationServiceClient(conReser),
		Restaurant:  restaurant.NewRestaurantServiceClient(conReser),
		Menu:        menu.NewMenuServiceClient(conReser),
	}
}
