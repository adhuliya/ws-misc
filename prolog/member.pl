mymember(X,[X|_]).
mymember(X,[_|T]) :-
  mymember(X,T).
