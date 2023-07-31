

intList (n :: Int)
    | n == 0 = []
    | otherwise = mod n 10 : intList(div n 10)


isPar n = l == reverse l
    where l = intList n
p4 = maximum [u | a <- [100..999], b <- [a..999], let u = a*b, isPar u]

p4' =  maximum [x | y<-[100..999], z<-[y..999], let x=y*z, let s=show x, s==reverse s]

main = do
    print p4
    print p4'