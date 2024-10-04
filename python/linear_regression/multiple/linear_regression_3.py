# Predicting housing prices based on sqft and # of bedrooms 
import numpy as np

# dataset
housing_data = np.array([[1800, 3], [2400, 4], [1416, 2], [3000, 5]])
prices = np.array([350000, 475000, 230000, 640000])

# Add 1's to matrix
X = np.append(np.ones(shape=(len(housing_data), 1)), housing_data, axis=1)

# Calculate matrix of coefficients, beta, using the Normal Equation
coefficients = np.linalg.inv(X.T @ X) @ X.T @ prices

# predicting prices
predicted_prices = X @ coefficients

# calculating residuals
residuals = prices - predicted_prices

# calculating total sum of squares
ss_total = np.sum(np.square(prices - np.mean(prices)))

# calculating residual sum of squares
ss_residuals = np.sum(np.square(residuals))

# calculating coefficient of determination
r2_score = 1 - (ss_residuals / ss_total)

print("Coefficients:", coefficients)
print("Residuals:", residuals)
print("Predicted Prices:", predicted_prices)
print("R^2 Score:", r2_score)