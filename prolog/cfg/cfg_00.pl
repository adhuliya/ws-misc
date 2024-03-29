% A simple cfg written in prolog.

% s1: a = 20
% s2: x = &a
% s3: print(*x)

%%% start

pred(start, start).

succ(start, s1).


%%% s1: a = 20

pred(s1, start).

def(s1).
def(s1, 'a').
type(s1, lhs, scalar('Int')).
type(s1, 'a', scalar('Int')).

expr(s1, lit).
lit('20').
type(s1, rhs, scalar('Int')).
type(s1, lit('20'), scalar('Int')).

succ(s1, s2).


%%% s2: x = &a

pred(s2, s1).

def(s2).
def(s2, 'x').
type(s2, lhs, ptr(scalar('Int'))).
type(s2, 'x', ptr(scalar('Int'))).

addrof(s2, 'a').
type(s2, rhs, ptr(scalar('Int'))).
type(s2, 'a', scalar('Int')).

succ(s2, s3).


%%% s3: print(*x)

pred(s3, s2).

nondef(s3).
dref(s3, 'x').
use(s3, 'x').
type(s3, 'x', ptr(scalar('Int'))).
type(s3, 'x', ptr(scalar('Int'))).

use(s3, Y) :-
  dref(s3, X),
  inpointee(s3, X, Y).

succ(s3, end).

%%% end

pred(end, s3).

weakdef(end).

succ(end, end).



%%% START inpointee definition
inpointee(start, X, nil).

inpointee(S, X, Y) :-
  pred(S, P),
  type(P, lhs, scalar(_)),
  inpointee(P, X, Y).

inpointee(S, X, Y) :-
  pred(S, P),
  def(P, X),
  addrof(P, Y).

inpointee(S, X, Y) :-
  pred(S, P),
  nondef(P),
  inpointee(P, X, Y).

inpointee(S, X, Y) :-
  pred(S, P),
  def(P, Z),
  inpointee(P, X, Y),
  X \= Z.

inpointee(S, X, Y) :-
  pred(S, P),
  weakdef(P),
  inpointee(P, X, Y).

%%% END   inpointee definition


%%% START outlive definition

outlive(S, X) :-
    succ(S, SS),
    use(SS, X).

outlive(S, X) :-
    succ(S, SS),
    nondef(SS),
    outlive(SS, X).

outlive(S, X) :-
    succ(S, SS),
    weakdef(SS),
    def(SS, X),
    outlive(SS, X).
    
outlive(S, X) :-
    succ(S, SS),
    def(SS, Y),
    outlive(SS, X),
    X \= Y.

%%% END   outlive definition


%inpointee(start, X, nil).
%
%inpointee(S, X, Y) :-
%  pred(S, P),
%  type(P, lhs, scalar(_)),
%  inpointee(P, X, Y).
%
%inpointee(S, X, Y) :-
%  pred(S, P),
%  def(P, X),
%  addrOf(P, Y).
%
%inpointee(S, X, Y) :-
%  pred(S, P),
%  weakdef(P),
%  inpointee(P, X, Y).
%
%outlive(S, X) :-
%    succ(S, SS),
%    use(SS, X).
%
%outlive(S, X) :-
%    succ(S, SS),
%    weakdef(SS),
%    def(SS, X),
%    outlive(SS, X).
%    
%outlive(S, X) :-
%    succ(S, SS),
%    def(SS, Y),
%    outlive(SS, X),
%    X \= Y.
%


