# If switch const

## Constants

Numeric constants increment automatically from the first value. The first value must be initialized.

## If statements

Variables declared inside a if statement are only in scope until the end of the if. If they have been declared before, the value they hold after the conditional is the value they held before.

## Switch statements

switch cases break by default. The keyword `fallthrough` is necessary to continue to the next case.

## Array slices

Builtin function `make` allows to allocate an array. It returns the array created. Builtin function `copy` allows to copy an array into a destination.

## Interface

More details about interfaces https://go101.org/article/interface.html.