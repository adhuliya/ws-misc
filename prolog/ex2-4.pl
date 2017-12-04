% http://www.learnprolognow.org/lpnpage.php?pagetype=html&pageid=lpn-htmlse7
word(astante,  a,s,t,a,n,t,e).
word(astoria,  a,s,t,o,r,i,a).
word(baratto,  b,a,r,a,t,t,o).
word(cobalto,  c,o,b,a,l,t,o).
word(pistola,  p,i,s,t,o,l,a).
word(statale,  s,t,a,t,a,l,e). 

crossword(V1, V2, V3, H1, H2, H3) :-
    word(V1, _,H1_2,_,H2_2,_,H3_2,_),
    word(V2, _,H1_4,_,H2_4,_,H3_4,_),
    word(V3, _,H1_6,_,H2_6,_,H3_6,_),
    word(H1, _,H1_2,_,H1_4,_,H1_6,_),
    word(H2, _,H2_2,_,H2_4,_,H2_6,_),
    word(H3, _,H3_2,_,H3_4,_,H3_6,_).
