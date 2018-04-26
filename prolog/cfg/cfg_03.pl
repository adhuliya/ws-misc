% A simple cfg written in prolog.

% s1: x = &y
% s2: y = &z
% s3: z = &u
% s4: *z = y
% s5: z = y
% s6: y = &x
% s7: print(u)
% s8: print(x)

% CFG
%    s1
%    s2
%    s3
% s4    s5
%    s6
%    s7
%    s8

%%% start

pred(start, start).

succ(start, s1).


%%% s1: x = &y

pred(s1, start).

def(s1).
def(s1, 'x').
type(s1, lhs, ptr(scalar('Int'))).
type(s1, 'x', ptr(scalar('Int'))).

addrof(s1, 'y').
type(s1, rhs, ptr(scalar('Int'))).
type(s1, 'y', ptr(scalar('Int'))).

succ(s1, s2).


%%% s2: y = &z

pred(s2, s1).

def(s2).
def(s2, 'y').
type(s2, lhs, ptr(scalar('Int'))).
type(s2, 'y', ptr(scalar('Int'))).

addrof(s2, 'z').
type(s2, rhs, ptr(scalar('Int'))).
type(s2, 'z', ptr(scalar('Int'))).

succ(s2, s3).


%%% s3: z = &u

pred(s3, s2).

def(s3).
def(s3, 'z').
type(s3, lhs, ptr(scalar('Int'))).
type(s3, 'z', ptr(scalar('Int'))).

addrof(s3, 'u').
type(s3, rhs, ptr(scalar('Int'))).
type(s3, 'u', ptr(scalar('Int'))).

succ(s3, s4).
succ(s3, s5).


%%% s4: *z = y

pred(s4, s3).

def(s4).
use(s4, 'z').
dref(s4, 'z').
type(s4, lhs, ptr(scalar('Int'))).
type(s4, 'z', ptr(scalar('Int'))).

use(s4, 'y').
type(s4, rhs, ptr(scalar('Int'))).
type(s4, 'y', ptr(scalar('Int'))).

succ(s4, s6).

def(s4, X) :-
    dref(s4, Z),
    inpointee(s4, Z, X).

addrof(s4, X) :-
    inpointee(s4, 'y', X).

weakdef(s4) :-
    dref(s4, Z),
    inpointee(s4, Z, X),
    inpointee(s4, Z, Y),
    X \= Y.


%%% s5: z = y

pred(s5, s3).

def(s5).
def(s5, 'z').
type(s5, lhs, ptr(scalar('Int'))).
type(s5, 'z', ptr(scalar('Int'))).

use(s5, 'y').
type(s5, rhs, ptr(scalar('Int'))).
type(s5, 'y', ptr(scalar('Int'))).

succ(s5, s6).

addrof(s5, X) :-
    inpointee(s5, 'y', X).


%%% s6: y = &x

pred(s6, s4).
pred(s6, s5).

def(s6).
def(s6, 'y').
type(s6, lhs, ptr(scalar('Int'))).
type(s6, 'y', ptr(scalar('Int'))).

addrof(s6, 'x').
type(s6, rhs, ptr(scalar('Int'))).
type(s6, 'x', ptr(scalar('Int'))).

succ(s6, s7).


%%% s7: print(u)

pred(s7, s6).

nondef(s7).
use(s7, 'u').
type(s7, 'u', ptr(scalar('Int'))).

succ(s7, s8).


%%% s8: print(x)

pred(s8, s7).

nondef(s8).
use(s8, 'x').
type(s8, 'x', ptr(scalar('Int'))).

succ(s8, end).


%%% end

pred(end, s8).

succ(end, end).


%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%


%%% START inpointee definition

inpointee(start, X, nil).

inpointee(S, X, Y) :-
  pred(S, SP),
  type(SP, lhs, scalar(_)),
  inpointee(SP, X, Y).

inpointee(S, X, Y) :-
  pred(S, SP),
  def(SP, X),
  addrof(SP, Y).

inpointee(S, X, Y) :-
  pred(S, SP),
  (nondef(SP) ; weakdef(SP)),
  inpointee(SP, X, Y).

inpointee(S, X, Y) :-
  pred(S, SP),
  def(SP, Z),
  inpointee(SP, X, Y),
  X \= Z.

%%% END   inpointee definition


%%% START outlive definition

outlive(S, X) :-
    succ(S, SS),
    use(SS, X).

outlive(S, X) :-
    succ(S, SS),
    (nondef(SS) ; weakdef(SS)),
    outlive(SS, X).

outlive(S, X) :-
    succ(S, SS),
    def(SS, Y),
    outlive(SS, X),
    X \= Y.

%%% END   outlive definition


%%% START LFCPA

lfcpa(S, X, Y) :-
    inpointee(S, X, Y),
    pred(S, SP),
    outlive(SP, X).

%%% END   LFCPA


