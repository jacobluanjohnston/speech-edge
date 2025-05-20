import random


# Provides fake sensor for dev/CI


# Define a mock version of the `INA219` class, so that when
# your main script tries to use `INA219`, it will use this
# mock instead.
class INA219:
    def __init__(self, i2c):
        pass

    @staticmethod
    def voltage():
        return round(random.uniform(0.8, 1.2), 2)

    @staticmethod
    def current():
        return round(random.uniform(-1.0, 1.0), 2)
