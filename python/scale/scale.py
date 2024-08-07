# Happy scale
# mean vs median
import numpy
results = [5, 4, 3, 4, 5, 3, 2, 5, 3, 2, 1, 4, 5, 3, 4, 4, 5, 4, 2, 1,
4, 5, 4, 3, 2, 4, 4, 5, 4, 3, 2, 1]

sorted_results = sorted(results)
print(sorted_results)

print(numpy.mean(results))
print(numpy.median(results))
median_index = int(len(results)/2)-1 # base 0
print("Median index: ", median_index)
print("Median: ", sorted_results[median_index])