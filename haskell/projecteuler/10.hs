

isqrt (n :: Integer) = floor $ sqrt $ fromInteger $ toInteger n

primes = 2 : [n | n <- [3,5..], all (\p -> mod n p  /= 0) $ takeWhile(<= isqrt n) primes]
p10 k = sum . takeWhile (<=k) $ primes

main = do
  print $ p10 2000_000