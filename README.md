# Linear-stats

## Overview
This Go program analyzes numerical data from a specified file to calculate two important statistical metrics: the Linear Regression Line and the Pearson Correlation Coefficient. The input data represents points on a graph where the x-values are derived from line numbers and the y-values are the numbers provided in the file.

Features:

 File Reading: Reads numerical data from a user-specified file.

Statistical Analysis: Computes the Linear Regression Line and the Pearson Correlation Coefficient.

Formatted Output: Presents results with specified decimal precision for clarity.

Input Format:

The program expects a text file where each line contains a single numerical value, like this:

189

113

121

114

145

110

...

x-axis: Represents line numbers (0, 1, 2, ...).

y-axis: Represents the numerical values from the file.

Getting Started:
## Prerequisites:

Latest Go version installed on your machine.

Running the Program

```bash
git clone  https://learn.zone01kisumu.ke/git/hilaromondi/linear-stats.git
```

```bash
cd linear-starts

$ go run . data.txt
```
Replace data.txt with the path to your input data file.

Output

Upon execution, the program will print the results in the following format:

Linear Regression Line: `y = <slope>x + <intercept>`

Pearson Correlation Coefficient: `<value>`

Formatting Requirements:

Linear Regression Line coefficients: Displayed to 6 decimal places.
Pearson Correlation Coefficient: Displayed to 10 decimal places.

Example Output

For an input file containing:

189

113

121

114

145 

110

...

The output might look like this:

Linear Regression Line: y = -0.155556x + 126.888889

Pearson Correlation Coefficient: 0.9254500000

* How It Works:

Data Reading: The program reads the values from the specified file.

Statistical Calculations:
Computes the Linear Regression using least squares method.

Calculates the Pearson Correlation Coefficient to measure the linear correlation between the variables.

Output: Displays the calculated metrics in a formatted manner.

## Author
name : [hilaromondi](https://learn.zone01kisumu.ke/git/hilaromondi/linear-stats.git)

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributions

Contributions are welcome! If you have suggestions or improvements, please fork the repository and create a pull request. Your feedback and contributions are greatly appreciated!