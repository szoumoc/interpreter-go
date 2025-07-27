Writing An Interpreter In Go


for lexer(It will take source code as input and output the tokens that represent the source
code.),
    "in a production
    environment it makes sense to attach filenames and line numbers to tokens, to better track
    down lexing and parsing errors. So it would be better to initialize the lexer with an io.Reader
    and the filename. But since that would add more complexity we’re not here to handle, we’ll
    start small and just use a string and ignore filenames and line numbers."