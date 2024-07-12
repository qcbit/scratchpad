import numpy as np
import matplotlib.pyplot as plt

x = np.array([1, 2, 3, 4, 5])
y = np.array([2, 4, 5, 4, 5])

mean_x = np.mean(x)
mean_y = np.mean(y)
print("Mean of x is ", mean_x)
print("Mean of y is ", mean_y)
print("Sum x - mean_x is ", np.sum(x - mean_x))
print("Sum y - mean_y is ", np.sum(y - mean_y))

numerator = np.sum((x - mean_x) * (y - mean_y))
denominator = np.sum((x - mean_x) ** 2)
print("Numerator is ", numerator)
print("Denominator is ", denominator)

m = np.sum((x - mean_x) * (y - mean_y)) / np.sum((x - mean_x) ** 2 )
b = mean_y - m * mean_x

print(f"y = {m}x + {b}")

plt.scatter(x, y, color = "red", marker="o", s=30)
y_pred = m * x + b
plt.plot(x, y_pred, color="blue")
plt.xlabel("x")
plt.ylabel("y")
plt.title("Simple Linear Regression")
plt.show()