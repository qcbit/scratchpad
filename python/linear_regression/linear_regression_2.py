# linear_regression_2.py
# Predicting if extra hour of advertising could boost restaurant sales
import numpy as np
import matplotlib.pyplot as plt

advertising_costs = np.array([100, 200, 300, 400, 500])
sales = np.array([300, 500, 600, 700, 800])

mean_ad_cost = np.mean(advertising_costs)
mean_sales = np.mean(sales)
print("Mean of ad cost is ", mean_ad_cost)
print("Mean of sales is ", mean_sales)

numerator = np.sum((advertising_costs - mean_ad_cost) * (sales - mean_sales))
denominator = np.sum((advertising_costs - mean_ad_cost) ** 2)
print("Numerator is ", numerator)
print("Denominator is ", denominator)

m = np.sum((advertising_costs - mean_ad_cost) * (sales - mean_sales)) / np.sum((advertising_costs - mean_ad_cost) ** 2 )
b = mean_sales - m * mean_ad_cost

print(f"Model: Sales = {m:.2f} * Advertising Costs + {b:.2f}")

plt.scatter(advertising_costs, sales, color = "red", marker="o", s=30)
y_pred = b + m * advertising_costs
plt.plot(advertising_costs, y_pred, color="red")
plt.xlabel("Advertising Costs")
plt.ylabel("Sales")
plt.title("Advertising Costs vs. Sales")
plt.show()