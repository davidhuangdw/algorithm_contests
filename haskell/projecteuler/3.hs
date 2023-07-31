import Data.List

p3 (n :: Integer) = last $ filter (\k -> isPrim k && n `mod` k == 0) [2 .. isqrt n]
  where
    isqrt (n :: Integer) = floor $ sqrt $ fromInteger $ toInteger n
    isPrim (n :: Integer) = all (\x -> n `mod` x /= 0) [2 .. isqrt n]

prims = 2 : filter (null . tail . primFactors) [3, 5 ..]

primFactors (n :: Integer) = factor n prims
  where
    factor n (p : ps)
      | p * p > n = [n]
      | mod n p == 0 = p : factor (div n p) (p : ps)
      | otherwise = factor n ps

p3' = maximum . primFactors

main = do
  print $ take 100 prims
  let n = 600851475143
  print $ primFactors n
  print $ p3' n
  -- print $ p3 n
