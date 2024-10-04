import numpy as np

# dataset
# real estate market model
coefficients = np.array([5000, 300, -1000])

# Data for a new house: [1, sqft, age]
houses = np.array([
  [1, 2000, 10],
  [1, 1500, 15],
  [1, 1800, 5]
  ])

# Use coefficients to predict the price of the new house
#predictions = houses.dot(coefficients) # equivalent to np.dot(houses, coefficients)
predictions = np.dot(houses, coefficients) # [595000 440000 540000]
print(predictions)