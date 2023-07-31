package google_code_jam.`2017`.`1a`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, P) = input_ints()
        val R = input_ints()
        val Q = (1..N).map{input_ints()}
        fun set2(): Int {
            var cnt = 0
            val A = Q.zip(R).map{ (_q, r) ->
                _q.map{q ->
                    (q*10 + 11*r-1)/(11*r) to (q*10)/(9*r)
                }.filter{(a, b) -> a <= b}.sortedWith(compareBy({it.first},{it.second}))  // bug: might be [a, x], [a, y] -- firsts are the same!!
            }
            val ind = MutableList(N){0}
            fun overlap(ps: List<Pair<Int, Int>>): Boolean{
                val l = ps.maxOf{it.first}
                val r = ps.minOf{it.second}
                return l <= r
            }
            while((0 until N).all{i -> ind[i] < A[i].size}){
                val intervals = ind.withIndex().map{(i, k) ->
                    A[i][k]
                }
                if(overlap(intervals)){
                    cnt ++
                    for(i in 0 until N){
                        ind[i] ++
                    }
                }else{
                    val i = (0 until N).minByOrNull { intervals[it].second}!!
                    ind[i] ++
                }
            }
            return cnt
        }
        val r = set2()
        println("Case #${it}: $r")
    }
}
