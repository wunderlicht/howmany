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
iterations? Run a lot of above scenarios and make a statistic!
I rolled the dice for 100,000 scenarios and put the results in the following
table. 😉

```
#iterations occurrence probably cumulative
          1          0     0.00       0.00
          2      28047    28.05      28.05
          3      46098    46.10      74.14
          4      20446    20.45      94.59
          5       4681     4.68      99.27
          6        672     0.67      99.94
          7         51     0.05      99.99
          8          4     0.00     100.00
          9          1     0.00     100.00
```

Of course `howmany` did the work for me.

Before looking at the table
* Were you aware that in no scenario it was possible to finish the job in 1
iteration?
* Did you see that commiting to do the job in 2 iterations is futile as
it only has a less than 30% chance?
* Did you see that committing to 3 iterations for the job is still likely to fail
in 1 out of 4 cases?

All of a sudden we can argue with probabilities on how many iterations we
should assume for a bunch of tasks at hand. We can adjust the planned iterations
based on the risk we are comfortable to take. When we need to hit a legal
deadline we should start at least 7 iterations in advance to be pretty sure
we will make it. If you ask me to finish some work for a fair that is after 2
iterations, I would politely decline because of the high risk NOT to finish on
time.

## How does it work?
`howmany` runs a lot of single scenarios based on historic data to answer one
simple question. How many iterations will it likely take?

A single scenario randomly picks historic data points and counts how many are
needed to hit (or overshoot) the given goal (9 in the above example). The
results of the scenarios are counted. Probabilities and cumulated probabilities
are calculated based on the resulting dataset of those simulations.

## Usage
`howmany -goal 55 -file history.csv -scenarios 300000 -confidence 90`  
same as  
`howmany -g 55 -f history.csv -s 300000 -c 90`

`-goal` and `-file` are mandatory parameters.

When `-scenarios` is omitted a default of 100,000 is assumed.  
When `-confidence` is omitted the respective marker is not printed.

All the above but goal can be set as defaults in environment variables
where the command line arguments take precedence. `HMFILE`, `HMSCENARIOS`,
`HMCONFIDENCE`

## Limitations
When you have historic data with a lot of zeros, a.k.a. nothing was delivered
during iterations, the single scenario might not terminate (the case is very
unlikely, still not impossible, so be aware). No measures were taken to catch
this highly unlikely case.