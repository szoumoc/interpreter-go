This project is based on the book Writing **An Interpreter In Go by Thorsten Ball**. It follows the development of a simple interpreter for a programming language, built from scratch using Go.

**Goal**
The aim is to implement a fully working interpreter that includes:

**A lexer** (tokenizer)

**A parser**

**An AST (Abstract Syntax Tree)**

**An evaluator**

Each component is built incrementally and explained thoroughly, making it an ideal resource for learning how interpreters work under the hood.

Lexer Design
The lexer is responsible for converting raw source code into a stream of tokens that the parser can understand.

Note from the book:
In a production environment, it makes sense to attach filenames and line numbers to tokens to better track down lexing and parsing errors. A more robust design would initialize the lexer with an io.Reader and a filename.

However, for the purposes of this learning-focused implementation, we've opted for a simpler approach: the lexer takes a raw string as input and ignores filenames and line numbers. This reduces complexity and keeps the focus on fundamental concepts.

This tradeoff allows us to move forward quickly and build a working interpreter, while leaving room for future improvements.

=======
**Writing An Interpreter In Go**


for **lexer(It will take source code as input and output the tokens that represent the source
code.**),
    "in a production
    environment it makes sense to attach filenames and line numbers to tokens, to better track
    down lexing and parsing errors. So it would be better to initialize the lexer with an io.Reader
    and the filename. But since that would add more complexity we’re not here to handle, we’ll
    start small and just use a string and ignore filenames and line numbers."

**Start of a REPL**

The Monkey language needs a REPL. REPL stands for **“Read Eval Print Loop”** and you probably know what it is from other interpreted languages: Python has a REPL, Ruby has one, every
JavaScript runtime has one, most Lisps have one and a lot of other languages too. Sometimes
the REPL is called “console”, sometimes “interactive mode”. The concept is the same: the
REPL reads input, sends it to the interpreter for evaluation, prints the result/output of the
interpreter and starts again. Read, Eval, Print, Loop.


**PARSERS**

A parser is a software component that takes input data (frequently text) and builds
a data structure – often some kind of parse tree, abstract syntax tree or other
hierarchical structure – giving a structural representation of the input, checking for
correct syntax in the process. […] The parser is often preceded by a separate lexical
analyser, which creates tokens from the sequence of input characters;
They take source code as input (either as text or tokens) and produce
a data structure which represents this source code. While building up the data structure, they
unavoidably analyse the input, checking that it conforms to the expected structure. Thus the
process of parsing is also called syntactic analysis.
A small example should make things clearer. Let’s say that we have the following source code:
`if (3 * 5 > 10) {
return "hello";
} else {
return "goodbye";
}`
And let’s say we are using JavaScript, have a MagicLexer, a MagicParser and the AST is built
out of JavaScript objects, then the parsing step might produce something like this:
`> var input = 'if (3 * 5 > 10) { return "hello"; } else { return "goodbye"; }';
> var tokens = MagicLexer.parse(input);
> MagicParser.parse(tokens);
{
type: "if-statement",
    condition: {
    type: "operator-expression",
    operator: ">",
    left: {
    type: "operator-expression",
    operator: "*",
    left: { type: "integer-literal", value: 3 },
    right: { type: "integer-literal", value: 5 }
    29
    },
    right: { type: "integer-literal", value: 10 }
    },
    consequence: {
    type: "return-statement",
    returnValue: { type: "string-literal", value: "hello" }
    },
    alternative: {
    type: "return-statement",
    returnValue: { type: "string-literal", value: "goodbye" }
    }
}`
The parser we are going to write is a recursive descent parser. And in particular, it’s a “top
down operator precedence” parser, sometimes called “Pratt parser”, after its inventor Vaughan
Pratt.
