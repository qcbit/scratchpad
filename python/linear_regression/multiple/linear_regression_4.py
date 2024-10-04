# Predicting house prices given the dataset
import numpy as np

# dataset
# Constants for area (m^2), number of bedrooms, and age of the house (year)
features = np.array([[120, 3, 10], [150, 4, 15], [90, 2, 20], [110, 3, 5]], dtype='float32')

# Corresponding house prices (in thousands of dollars)
prices = np.array([[300], [400], [250], [275]])

# Add 1's to matrix
X = np.append(np.ones(shape=(len(features), 1)), features, axis=1)

# Calculate matrix of coefficients, beta, using the Normal Equation
coefficients = np.linalg.inv(X.T @ X) @ X.T @ prices

# predicting prices
predicted_prices = X @ coefficients

# calculating residuals
residuals = prices - predicted_prices

# calculating total sum of squares
ss_total = ((prices - prices.mean()) ** 2).sum()

# calculating residual sum of squares
ss_residuals = ((prices - predicted_prices) ** 2).sum()

# calculating coefficient of determination
r2_score = 1 - ss_residuals / ss_total

print("Coefficients:", coefficients)
print("Residuals:", residuals)
print("Predicted Prices:", predicted_prices)
print("R^2 Score:", r2_score)