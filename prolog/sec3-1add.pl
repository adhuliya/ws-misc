add(0,X,Z) :- X = Z.
add(succ(Y),X,Z) :- add(Y,succ(X),Z).
