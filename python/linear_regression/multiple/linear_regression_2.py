import numpy as np

# dataset
# real estate market model
coefficients = np.array([50000, 3000, -2000, 15000])

# Data for a new house: [intercept, area, age, # of bedrooms]
new_house = np.array([1, 150, 10, 2])

# Use coefficients to predict the price of the new house
predictions = new_house.dot(coefficients)
print(predictions)