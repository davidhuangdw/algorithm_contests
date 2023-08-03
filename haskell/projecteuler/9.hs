
p9 = [(a, b, c) | a<-[1..1000], b<-[a..1000-a], let c = 1000-a-b, b < c, a * a + b * b == c * c]

main = do
  print p9