append2([],X,X).
append2([H|T],X,[H|T1]) :-
    append2(T,X,T1).


