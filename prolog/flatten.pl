flatten([],[]) :- !.
flatten([H|T], Y) :-
    !,
    flatten(H, Y1),
    flatten(T, Y2),
    append(Y1, Y2, Y).

flatten(X,[X]).

