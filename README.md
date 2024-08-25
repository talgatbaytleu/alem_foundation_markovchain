# Welcome to tbaitleu/marcov-chain git page

By default this program reads from stdin the whole text and generate the text according to Markov Chain algorithm.

Outcomes:

+ Program prints generated text according to the Markov Chain algorithm.

Notes:

+ Suffix length is ALWAYS 1 word
+ Default prefix length is 2 words.
+ Default starting prefix is the first N words of the text, where N is the length of the prefix.
+ Default number of maximum words is 100.

## Here is the little instraction how to use this text generator

### 1. Firstly please build executible file if it haven't been builded yet.

``go build build -o markovchain .``

### 2. Since it recieves dictionary data from stdin, You'll need to pipe the content of file where you store your dictory data:

For example:

``cat the-great-gatsby.txt | ./markovchain``

### 3. It has three values 
