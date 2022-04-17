package google_code_jam.`2020`.`1b`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (R, S) = input_ints()
        fun set1(): String{
            val n = R * S
            val st = (1..S).flatMap{ (1..R).map{ it }}
            val ed = (1..R).flatMap{ r -> (1..S).map{r}}
            val done = HashMap<List<Int>, Triple<Int, Int, List<Int>>>()
            val que = ArrayDeque<Pair<List<Int>, Int>>()
            done[ed] = Triple(-1, -1, listOf())
            que.add(ed to 0)
            while(que.isNotEmpty()){
                var (s, k) = que.removeFirst()
                k ++
                for(a in 1 until n-1)
                    for(b in 1..n-a){
                        val v = s.subList(a, a+b) + s.subList(0, a)  + s.subList(a+b, n)
                        if(v == st) {
                            val moves = mutableListOf<String>()
                            var (x, y) = a to b
                            var pre = s
                            while(x != -1){
                                moves.add("$y $x\n")
                                val dn = done[pre]!!
                                x = dn.first
                                y = dn.second
                                pre = dn.third
                            }
                            return buildString {
                                append("${moves.size}\n")
                                for(m in moves){
                                    append(m)
                                }
                            }
                        }
                        if(v in done) continue
                        done[v] = Triple(a, b, s)
                        que.addLast(v to k)
                    }
            }
            throw RuntimeException("not found")
        }
        val res = set1()
        // todo: set2
        print("Case #${it}: $res")
    }
}
