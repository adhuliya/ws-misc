-- A haskell code
-- Quicksort in Haskell
qsort (x:xs) = qsort left ++ [x] ++ qsort right
  where
    left = filter (<x) xs
    right = filter (>=x) xs
