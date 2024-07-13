# linear_regression_3.py
# Predicting restaurant sales if extra hour of advertising could boost restaurant sales
import numpy as np
import matplotlib.pyplot as plt

ad_hours = np.array([10, 20, 30, 40, 50])
weekly_sales = np.array([200, 420, 650, 800, 950])

mean_ad_hours = np.mean(ad_hours)
mean_weekly_sales = np.mean(weekly_sales)
print("Mean of ad hours is ", mean_ad_hours)
print("Mean of weekly sales is ", mean_weekly_sales)

numerator = np.sum((ad_hours - mean_ad_hours) * (weekly_sales - mean_weekly_sales))
denominator = np.sum((ad_hours - mean_ad_hours) ** 2)
print("Numerator is ", numerator)
print("Denominator is ", denominator)

m = np.sum((ad_hours - mean_ad_hours) * (weekly_sales - mean_weekly_sales)) / np.sum((ad_hours - mean_ad_hours) ** 2 )
b = mean_weekly_sales - m * mean_ad_hours

print(f"Model: Sales = {m:.2f} * Ad_Hours + {b:.2f}")

plt.scatter(ad_hours, weekly_sales, color = "red", marker="o", s=30)
y_pred = b + m * ad_hours
plt.plot(ad_hours, y_pred, color="red")
plt.xlabel("Ad Hours")
plt.ylabel("Weekly Sales")
plt.title("Ad Hours vs. Weekly Sales")
plt.show()