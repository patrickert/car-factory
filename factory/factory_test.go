package factory

import (
	".main.go/vehicle"
	"testing"

	"github.com/stretchr/testify/suite"
)

type factoryUnitTestSuite struct {
	suite.Suite
	adapter *Factory
}

func (s *factoryUnitTestSuite) SetupSuite() {

	s.adapter = &Factory{}
}

func TestFactoryUnitTestSuite(t *testing.T) {
	suite.Run(t, &factoryUnitTestSuite{})
}

func (s *factoryUnitTestSuite) TestSamble() {
	carsAmount := 10
	factory := New()
	ch := make(chan vehicle.Car, carsAmount)
	factory.StartAssemblingProcess(carsAmount, ch)

	expectedCar := vehicle.Car{
		Id:            0,
		Chassis:       "Assembled",
		Tires:         "Assembled",
		Engine:        "Assembled",
		Electronics:   "Assembled",
		Dash:          "Assembled",
		Sits:          "Assembled",
		Windows:       "Assembled",
		EngineStarted: true,
		TestingLog:    "Engine Started!, Moved forward 10 meters!, Moved forward 10 meters!, Turned Right, Turned Right!, Engine Stopped!, ",
		AssembleLog:   "Assembled",
	}
	for i := 0; i < carsAmount; i++ {
		car := <-ch
		s.Assert().Equal(expectedCar.Chassis, car.Chassis)
		s.Assert().Equal(expectedCar.Tires, car.Tires)
		s.Assert().Equal(expectedCar.Engine, car.Engine)
		s.Assert().Equal(expectedCar.Electronics, car.Electronics)
		s.Assert().Equal(expectedCar.Dash, car.Dash)
		s.Assert().Equal(expectedCar.Sits, car.Sits)
		s.Assert().Equal(expectedCar.Windows, car.Windows)
		s.Assert().Equal(expectedCar.EngineStarted, car.EngineStarted)
		s.Assert().Equal(expectedCar.TestingLog, car.TestingLog)
		s.Assert().NotEmpty(expectedCar.AssembleLog)
	}

	//code here
	// Assert
	s.Assert().Equal(1, 1)
}
