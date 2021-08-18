package assemblyspot

import (
	"errors"
	"fmt"
	"sync"
	"time"

	".main.go/vehicle"
)

type AssemblySpot struct {
	vehicleToAssemble *vehicle.Car
	assemblyLog       string
	logMu             *sync.Mutex
}

func New() *AssemblySpot {
	return &AssemblySpot{
		vehicleToAssemble: nil,
		assemblyLog:       "",
		logMu:             &sync.Mutex{},
	}
}

func (s *AssemblySpot) SetVehicle(v *vehicle.Car) {
	s.vehicleToAssemble = v
}

func (s *AssemblySpot) GetAssembledVehicle() *vehicle.Car {
	return s.vehicleToAssemble
}

func (s *AssemblySpot) IsAvailable() bool {
	return s.vehicleToAssemble == nil
}

func (s *AssemblySpot) GetAssembledLogs() string {
	return s.assemblyLog
}

//hint: improve this function to execute this process concurrenlty
func (s *AssemblySpot) AssembleVehicle() (*vehicle.Car, error) {
	if s.vehicleToAssemble == nil {
		return nil, errors.New("no vehicle set to start assembling")
	}

	wg := &sync.WaitGroup{}
	wg.Add(7)

	runConcurrently(s.assembleChassis, wg)
	runConcurrently(s.assembleTires, wg)
	runConcurrently(s.assembleEngine, wg)
	runConcurrently(s.assembleElectronics, wg)
	runConcurrently(s.assembleDash, wg)
	runConcurrently(s.assembleSeats, wg)
	runConcurrently(s.assembleWindows, wg)

	wg.Wait()
	return s.vehicleToAssemble, nil
}

func (s *AssemblySpot) assembleChassis() {
	s.vehicleToAssemble.Chassis = "Assembled"
	time.Sleep(1 * time.Second)
	s.safeLogUpdate(fmt.Sprintf("Chassis at [%s], ", time.Now().Format("2006-01-02 15:04:05.000")))
}

func (s *AssemblySpot) assembleTires() {
	s.vehicleToAssemble.Tires = "Assembled"
	time.Sleep(1 * time.Second)
	s.safeLogUpdate(fmt.Sprintf("Tires at [%s], ", time.Now().Format("2006-01-02 15:04:05.000")))
}

func (s *AssemblySpot) assembleEngine() {
	s.vehicleToAssemble.Engine = "Assembled"
	time.Sleep(1 * time.Second)
	s.safeLogUpdate(fmt.Sprintf("Engine at [%s], ", time.Now().Format("2006-01-02 15:04:05.000")))
}

func (s *AssemblySpot) assembleElectronics() {
	s.vehicleToAssemble.Electronics = "Assembled"
	time.Sleep(1 * time.Second)
	s.safeLogUpdate(fmt.Sprintf("Electronics at [%s], ", time.Now().Format("2006-01-02 15:04:05.000")))
}

func (s *AssemblySpot) assembleDash() {
	s.vehicleToAssemble.Dash = "Assembled"
	time.Sleep(1 * time.Second)
	s.safeLogUpdate(fmt.Sprintf("Dash at [%s], ", time.Now().Format("2006-01-02 15:04:05.000")))
}

func (s *AssemblySpot) assembleSeats() {
	s.vehicleToAssemble.Sits = "Assembled"
	time.Sleep(1 * time.Second)
	s.safeLogUpdate(fmt.Sprintf("Sits at [%s], ", time.Now().Format("2006-01-02 15:04:05.000")))
}

func (s *AssemblySpot) assembleWindows() {
	s.vehicleToAssemble.Windows = "Assembled"
	time.Sleep(1 * time.Second)
	s.safeLogUpdate(fmt.Sprintf("Windows at [%s], ", time.Now().Format("2006-01-02 15:04:05.000")))
}

func (s *AssemblySpot) safeLogUpdate(log string) {
	s.logMu.Lock()
	s.assemblyLog += log
	s.logMu.Unlock()
}

func (s *AssemblySpot) ClearLog() {
	s.assemblyLog = ""
}

func runConcurrently(f func(), group *sync.WaitGroup) {
	go func() {
		f()
		group.Done()
	}()
}
