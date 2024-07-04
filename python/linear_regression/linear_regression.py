import numpy as np
import matplotlib.pyplot as plt

x = np.array([1, 2, 3, 4, 5])
y = np.array([2, 4, 5, 4, 5])

mean_x = np.mean(x)
mean_y = np.mean(y)

m = np.sum((x - mean_x) * (y - mean_y)) / np.sum((x - mean_x) ** 2 )
c = mean_y - m * mean_x

print(f"y = {m}x + {c}")

plt.scatter(x, y, color = "red", marker="o", s=30)
y_pred = m * x + c
plt.plot(x, y_pred, color="blue")
plt.xlabel("x")
plt.ylabel("y")
plt.title("Simple Linear Regression")
plt.show()