# AoC Reporter (work-in-progress)

**CLI reporter of Advent of Code data** 

- "How many gold stars were earned in total this year?"
- "What percentage of people completed both parts (vs. only part 1) of puzzle 14 this year?"
- "..."

Questions like this and more can be answered using AoC Reporter.

## Flags

- `-s`, the statistics to be reported. Currently available values are `total`
  and `graph`. If not specified, a default value of `graph` is used.
- `-y`, the year to be reported. If not specified, a default value of `2023` is
  used.
