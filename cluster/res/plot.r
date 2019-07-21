x1 <- c(-10, -10, -8, -8, 10, 10, 8, 8);
x2 <- c(-10, -8, -8, -10, 10, 8, 8, 10);

plot(x1, x2);

points(c(-9, 9), c(-9, 9), col='purple');

library(plotrix)
draw.circle(6.8, 7.5, sqrt(31.5), border="red");
draw.circle(-0.3, -0.3, sqrt(81.5), border="blue");
