

p1 = sum [x | x <- [1..999], x `mod` 3 == 0 || x `mod` 5 == 0]

p1' = asum 3 + asum 5 - asum 15
    where asum k = k * acc (999 `div` k)
          acc k = (1+k) * k `div` 2

main = do
    print p1
    print p1'