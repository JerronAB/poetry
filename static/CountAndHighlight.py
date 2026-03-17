from js import console
import sys
from random import randint, shuffle
sys.stdout.write = lambda s: console.log(s)
#Highlight words only if they appear in more
#than one poem?
def main():
    exclude={'','am','what','like','me','through','when','where','my','a','i','an','and','are','as','at','be','but','by','for','if','in','into','is','it','no','not','of','on','or','such','that','the','their','then','there','these','they','this','to','was','will','with','from','you','your','i','we','he','she','his','her','them','our','us','so','too','have','has','had'}
    t = poems_content.lower().split()
    t = tuple(word for word in t if word not in exclude)
    counter_dict = {w:0 for w in t}
    for word in t: counter_dict[word] += 1
    print(dict(sorted(counter_dict.items(), key=lambda kv: kv[1], reverse=True)))
    counts = sorted([v for v in counter_dict.values() if v > 1], reverse=True)
    elig_counts = set(counts[:4])
    print(f'Words of this frequency are eligible: {elig_counts}')
    eligible_words = [word for word, count in counter_dict.items() if count in elig_counts]
    print(eligible_words)
    display_n_words = randint(1,3) #display 1-3 words
    shuffle(eligible_words)
    return eligible_words[:display_n_words]
    
    

#keeping this here allows the 
#returned value to implicitly become a
#javascript array
main()