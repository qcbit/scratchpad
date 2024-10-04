# Standard Deviation
import numpy

# data
temps = [31, 32, 32, 31, 28, 29, 31, 38, 32, 31, 30, 29, 30, 31, 26]
print("Temps: ", temps)

# 1. Find the mean
mean = numpy.mean(temps)
print("Mean: ", mean)

# 2. For each dataset, subtract the mean and square it
squared_differences = [(temp - mean)**2 for temp in temps]

# 3. Find the average of each squared difference
mean_of_squared_differences = numpy.mean(squared_differences)
print("Mean of squared differences: ", mean_of_squared_differences)

# 4. Find the standard deviation by taking the square root of the squared difference
standard_deviation = numpy.sqrt(mean_of_squared_differences)

print("Standard deviation: ", standard_deviation)