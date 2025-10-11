+++
draft = true
math = true
Toc = true
+++

# MATH
This is an inline \(a^*=x-b^*\) equation.

These are block equations:

\[a^*=x-b^*\]

\[ a^*=x-b^* \]

\[
a^*=x-b^*
\]

These are also block equations:

$$a^*=x-b^*$$

$$ a^*=x-b^* $$

$$
a^*=x-b^*
$$

# Direct HTML
<div class="framed pretty">
<p>
for those who are already semi-familiar with how word embeddings work, if you just want to skip to the meet where I actually implement in Excel, you can collapse the explanations below.
</p>
</div>

# Expandable Details

{{< details summary="Why Excel? (click to expand)" class="framed" open="false" >}}
Of course, Excel isn't my first choice for a process like this, but I'm limited by some constraints in my organization. firstly, I'd like to build a process that lives beyond me, and most folks sitting at a desk know how to at least open and poke around in Excel. if I wrapped this process up with python or another language of my choice, it would become completely unmaintainable in the future, since we can't really expect anyone who gets this task to be skilled in the language. 
secondly, it's actually pretty educational to see the sort of visual aid that Excel naturally provides for the high dimensional representation of a word. But, more to come on that. 
{{< /details >}}

# Images
{{< img src="vactubes.png" >}}
or
![optional description](/kmap_adjacency.png)

{{< asterisk class="pretty framed" symbol="*" >}}
If this specific kind of solution interests you, check out Markov chains and ... . these are other "classical" used to make language or perform language analysis, that don't get talked about as much these days. 
{{< /asterisk >}}