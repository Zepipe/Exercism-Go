package speed
// Car представляет машину с дистанционным управлением
type Car struct {
    battery      int // текущий заряд батареи в процентах
    batteryDrain int // расход батареи за одну поездку в процентах
    speed        int // скорость в метрах за одну поездку
    distance     int // пройденная дистанция в метрах
}

// Track представляет гоночную трассу
type Track struct {
    distance int // длина трассы в метрах
}

// NewCar создает новую машину с заданной скоростью и расходом батареи
func NewCar(speed, batteryDrain int) Car {
    return Car{
        speed:        speed,
        batteryDrain: batteryDrain,
        battery:      100, // начальный заряд 100%
        distance:     0,   // начальная дистанция 0
    }
}

// NewTrack создает новую трассу с заданной дистанцией
func NewTrack(distance int) Track {
    return Track{
        distance: distance,
    }
}

// Drive управляет машиной один раз
func Drive(car Car) Car {
    // Проверяем, достаточно ли батареи для поездки
    if car.battery < car.batteryDrain {
        // Недостаточно батареи - машина не двигается
        return car
    }
    
    // Обновляем состояние машины
    return Car{
        speed:        car.speed,
        batteryDrain: car.batteryDrain,
        battery:      car.battery - car.batteryDrain, // уменьшаем заряд
        distance:     car.distance + car.speed,        // увеличиваем дистанцию
    }
}

// CanFinish проверяет, может ли машина завершить гонку
func CanFinish(car Car, track Track) bool {
    // Рассчитываем, сколько раз машина может проехать с текущим зарядом
    maxDrives := car.battery / car.batteryDrain
    
    // Максимальная дистанция, которую может проехать машина
    maxDistance := maxDrives * car.speed
    
    // Проверяем, достаточно ли максимальной дистанции для трассы
    return maxDistance >= track.distance
}