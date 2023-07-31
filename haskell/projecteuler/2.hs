


p2 = sum [x| x <- xs, even x && x <= 1000_000]
    where   xs = fibs 100
            fibs 1 = [1]
            fibs 2 = [1, 1]
            fibs n = a+b:xs
                where   xs = fibs(n-1)
                        a:b:_ = xs

p2' = sum $ filter even $ takeWhile (<= 1000_000) fibs
    where fibs = 1 : 1 : zipWith (+) fibs (tail fibs)

main = do
    print p2
    print p2'