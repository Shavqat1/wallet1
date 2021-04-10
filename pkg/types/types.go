package types

//Money представляет собой денежную сумму в минимальных удиницах (центы, копейки, дирамы, и т.д.)
type Money int64

//PaymentCategory представляет собой категорию,в которой был совершён платёж (авто, аптеки, рестораны и т.д.)
type PaymentCategory string
//PaymentStatus представляет собой статус платежа.
type PaymentStatus string

//Предопределённые статусы плаьежей.
const (
	PaymentStatusOk PaymentStatus="Ok"
	PaymentStatusFail PaymentStatus="Fail"
	PaymentStatusInProgress PaymentStatus="INPROGRESS"
)
//Payment представляет инфрматцию о платеже.
type Payment struct {
	ID string
	Amount Money
	Category PaymentCategory
	Status PaymentStatus
}
type Phone string

//Account  представляет информацию о счёте пользователя.
type Account struct {
	ID int
	Phone Phone
	Balance Money
}
