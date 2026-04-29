package carservice

import "fmt"

type Appointment struct {
	CarNumber   string
	OwnerName   string
	ServiceType string
}

type AppointmentManager struct {
	appointments []Appointment
}

var validServices = []string{"ТО", "Ремонт", "Диагностика"}

func (m *AppointmentManager) CreateAppointment(carNumber, ownerName, serviceType string) error {
	if carNumber == "" {
		return fmt.Errorf("номер автомобиля не может быть пустым")
	}
	if ownerName == "" {
		return fmt.Errorf("имя владельца не может быть пустым")
	}

	validService := false
	for _, s := range validServices {
		if s == serviceType {
			validService = true
			break
		}
	}
	if !validService {
		return fmt.Errorf("неизвестный тип услуги: %s", serviceType)
	}

	for _, a := range m.appointments {
		if a.CarNumber == carNumber {
			return fmt.Errorf("запись с номером %s уже существует", carNumber)
		}
	}

	m.appointments = append(m.appointments, Appointment{
		CarNumber:   carNumber,
		OwnerName:   ownerName,
		ServiceType: serviceType,
	})
	return nil
}

func (m *AppointmentManager) FindAppointment(carNumber string) (Appointment, error) {
	if carNumber == "" {
		return Appointment{}, fmt.Errorf("номер автомобиля не может быть пустым")
	}

	for _, a := range m.appointments {
		if a.CarNumber == carNumber {
			return a, nil
		}
	}
	return Appointment{}, fmt.Errorf("запись с номером %s не найдена", carNumber)
}

func (m *AppointmentManager) CancelAppointment(carNumber string) error {
	if carNumber == "" {
		return fmt.Errorf("номер автомобиля не может быть пустым")
	}

	for i, a := range m.appointments {
		if a.CarNumber == carNumber {
			m.appointments = append(m.appointments[:i], m.appointments[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("запись с номером %s не найдена", carNumber)
}