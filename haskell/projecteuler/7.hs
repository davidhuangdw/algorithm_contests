
primes = 2 : [n | n <- [3,5..], all (\p -> mod n p  /= 0) $ takeWhile(\p -> p*p <= n) primes]
p7 = primes !! 10_000

main = do
  print p7
