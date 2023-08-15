import Data.Array
import Data.Foldable (maximumBy)
import Data.Ord (comparing)

collatz n
  | n == 1 = [1]
  | otherwise = n : collatz nxt
  where
    nxt = if even n then div n 2 else 3 * n + 1

collatzs n = a
  where
    a = listArray (1, n) $ (1, [1]) : map sol [2 .. n]
    sol m = (cnt + 1, m : ls)
      where
        nxt = if even m then div m 2 else m * 3 + 1
        (cnt, ls) = if nxt < n then a ! nxt else sol nxt

problem14 = maximumBy (comparing fst) . collatzs $ 1000_000

main = do
  -- print $ collatz 1000
  let (cnt, chain) = problem14
  print (cnt, chain)
  print $ collatz (head chain)