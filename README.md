# howmany
How many (iterations)? A.k.a. when is it done?

## Why does it exist?
There was, is, will be the question when is it done.
Totally understandable and as software engineers we failed way to often
answering this question.
A formula used by my professor in the 90s was:

Length of a task = (gut feeling + 20%) * 2

Gut feeling of one week became 6 days * 2, 12 days, so 2 working weeks and
2 days. Of course this was far from the truth and the wisdom in the formula
is questionable.

Then the agile guys came up with time boxes (best ever), story point
estimations, and the velocity of a team. (definition of velocity: rolling
average of the last 8 sprints). This is slightly better, still misleading.

Here is a little experiment. Let's assume we have a (preliminary) velocity of
3.5. Conveniently the historic data was 1, 2, 3, 4, 5, 6 points per iteration
(so you can follow along with a normal dice). How many iterations would it take
to get 9 points done? It should take 9/3.5 ≈ 2.6, round up to 3, iterations.
Let's roll the dice!

![rolled dice 2,6,2](assets/262.webp)
![rolled dice 3,1,4,4](assets/3144.webp)
![rolled dice 4,5](assets/45.webp)
![rolled dice 2,2,1,2,3](assets/22123.webp)

We got 3, 4, 2, 5 iterations as results. Obviously 2.6 is really just an
estimate and the truth varies. This bears one question.
What is the probability of the amount of work done in 1, 2, 3, 4, 5,...
iterations? Run a lot of above scenarios and make a statistic! This is exactly
what `howmany` does.

## How does it work?
`howmany` runs a lot of single scenarios based on historic data to answer one
simple question. How many iterations will it likely take?