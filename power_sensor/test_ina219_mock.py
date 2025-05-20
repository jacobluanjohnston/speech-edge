from power_sensor.mock_ina219 import INA219


def test_mock_sensor_range():
    sensor = INA219(i2c=None)  # i2c is irrelevant
    voltage = sensor.voltage()
    current = sensor.current()

    assert 0.8 <= voltage <= 1.2
    assert -1.0 <= current <= 1.0
