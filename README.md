# Bruhlang
Bruhlang is ***bruh*** word based esoteric language. Syntax inspired by [Zwedgy/Bruh](https://github.com/Zwedgy/Bruh)

_**Build:**_ `go build .`

**_Usage:_** `./bruhlang <file>`

<hr>

## Syntax:

### Keywords
- bruh
- moment
- momentum
- effect
- sound 

| Bruhlang  | Meaning  |
|---|---|
| bruh \<variable>  | Define variable  |
| \<variable> momentum  | Increment variable by one  |
| \<variable> moment  | Decrement variable by one  |
| moment \<variable>  | Print variable as integer  |
| momentum \<variable>  | Print variable as ASCII  |
| \<variable> \<variable_two>  | Assign \<variable_two> to \<variable>  |
| \<variable> effect \<variable_two>  | Increment \<variable> by \<variable_two>  |
| \<variable> sound \<variable_two>  | Decrement \<variable> by \<variable_two>  |
| sound -\<variable> \<variable_two>- >> \<statement> << sound  | Loop \<statement> until \<variable> is equal to \<variable_two>  |
| sound -\<variable> momentum \<variable_two>- >> \<statement> << sound  | Loop \<statement> until \<variable> is bigger than \<variable_two>  |
| sound -\<variable> moment \<variable_two>- >> \<statement> << sound  | Loop \<statement> until \<variable> is smaller than \<variable_two>  |

**Note:** To get multiple statements on while, you can use ***|*** operator. Example: `sound -aa moment bb- >> moment aa | aa momentum << sound`

**Note 2:** You need to put ***||*** on end of lines.

## Example:
    bruh b ||  
    bruh bb ||  
    bruh h ||  
      
    b momentum ||  
    b momentum ||  
    b momentum ||  
    b momentum ||  
    b momentum ||  
    b momentum ||  
    b momentum ||  
    b momentum ||  
    b momentum ||  
    b momentum ||  
    b momentum ||  
    b effect b ||  
    bb b ||  
    b effect b ||  
    b effect bb ||  
    momentum b ||  
      
    b effect b ||  
    b sound bb ||  
    h b ||  
    b momentum ||  
    b momentum ||  
    b momentum ||  
    b momentum ||  
    momentum b ||  
      
    b momentum ||  
    b momentum ||  
    b momentum ||  
    momentum b ||  
      
    h moment ||  
    h moment ||  
    h moment ||  
    h moment ||  
    h moment ||  
    h moment ||  
    momentum h ||

**_Output_:** `Bruh`