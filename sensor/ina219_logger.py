import time
from adafruit_ina219 import INA219
from busio import I2C
from board import SCL, SDA

i2c = I2C(SCL, SDA)
sensor = INA219(i2c)

while True:
    print(f"Bus Voltage: {sensor.bus_voltage:.2f} V")
    print(f"Current: {sensor.current:.2f} mA")
    print("-----")
    time.sleep(2)
