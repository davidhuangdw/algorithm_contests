package google_code_jam.`2019`.`1a`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (R, C) = input_ints()
        fun set1(R: Int, C: Int): MutableList<Pair<Int, Int>>{
            val sum = R * C
            val path = mutableListOf<Pair<Int, Int>>()
            val done = hashSetOf<Pair<Int, Int>>()
            fun dfs(pi: Int, pj: Int): Boolean {
                if(path.size == R * C){
                    return true
                }
                for(i in 1..R)
                    for(j in 1..C){
                        if(i to j !in done && i != pi && j != pj && i+j != pi+pj && i-j != pi - pj){
                            path.add(i to j)
                            done.add(i to j)
                            if(dfs(i, j)) return true
                            done.remove(i to j)
                            path.removeLast()
                        }
                    }
                return false
            }
            dfs(100, 200)
            return path
        }

        fun set2(R: Int, C: Int): MutableList<Pair<Int, Int>>{ // greedy..
            val path = mutableListOf<Pair<Int, Int>>()
            val done = hashSetOf<Pair<Int, Int>>()
            val rem  = List(4){hashMapOf<Int, Int>()}
            fun add(i: Int, j: Int, v: Int){
                rem[0][i] = rem[0].getOrDefault(i, 0) + v
                rem[1][j] = rem[1].getOrDefault(j, 0) + v
                rem[2][i+j] = rem[2].getOrDefault(i+j, 0) + v
                rem[3][i-j] = rem[3].getOrDefault(i-j, 0) + v
            }
            for(i in 1..R)
                for(j in 1..C){
                    add(i, j, 1)
                }
            fun deg(i: Int, j: Int): Int = rem[0][i]!! + rem[1][j]!! + rem[2][i+j]!! + rem[3][i-j]!!
            var (pi, pj) = 100 to 200
            for(_i in 1..R*C){
                var (mi, mj) = -1 to -1
                var max = 0
                for(i in 1..R)
                    for(j in 1..C){
                        val p = i to j
                        if(p !in done && i != pi && j != pj && i+j != pi + pj && i-j != pi-pj){
                            val d = deg(i, j)
                            if(d > max){
                                max = d
                                mi = i
                                mj = j
                            }
                        }
                    }
                if(mi < 0)
                    return mutableListOf()
                done.add(mi to mj)
                path.add(mi to mj)
                add(mi, mj, -1)
                pi = mi
                pj = mj
            }
            return path
        }

        val path = set2(R, C)
        val res = if(path.size == 0) "IMPOSSIBLE" else buildString{
            append("POSSIBLE")
            for((i, j) in path){
                append("\n$i $j")
            }
        }
        println("Case #${it}: $res")
    }
}
