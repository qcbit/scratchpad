import numpy as np

# dataset
X = np.array([[73, 67, 43],
  [91, 88, 64],
  [87, 134, 58],
  [102, 43, 37],
  [69, 96, 70]], dtype='float32')

y = np.array([56, 81, 119, 22, 103], dtype='float32')

# Calculate matrix of coefficients, beta, using the Normal Equation
# Add a column of 1's to X
ones = np.ones(shape=(len(X), 1))
X = np.append(ones, X, axis=1)

# Compute coefficients
# beta = np.linalg.inv(X.T.dot(X)).dot(X.T).dot(y)
beta = np.linalg.inv(X.T @ X) @ X.T @ y

print(beta)

# Model's Performance Evaluation
# coefficient of determination to see how well model fits the data
predictions = X.dot(beta)
ss_residuals = np.sum(np.square(y - predictions))
ss_total = np.sum(np.square(y - np.mean(y)))
r2_score = 1 - (ss_residuals / ss_total)
print("R^2 Score:", r2_score)