import Data.List

primes = 2 : filter isPrime [3, 5 ..]
  where
    isPrime = not . null . pFactors
-- isPrime n = all ((/= 0) . mod n) $ takeWhile ((<= n) . (^ 2)) primes

pFactors n = go n primes
  where
    go m (p : ps)
      | p * p > m = [m]
      | mod m p == 0 = p:go (div m p) (p : ps)
      | otherwise = go m ps

dividorNum = product . map ((+ 1) . length) . group . pFactors

triangulars = scanl1 (+) [1 ..]

problem12 = head . filter ((>= 500) . dividorNum) $ triangulars

main = do
  -- print $ take 100 primes
  -- print $ factorNums 28
  -- print $ take 100 triangulars
  let n = problem12
  print (n, pFactors n, dividorNum n)
