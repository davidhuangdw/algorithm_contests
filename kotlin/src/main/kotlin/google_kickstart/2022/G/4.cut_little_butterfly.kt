package google_kickstart.`2022`.G
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, E) = input_ints()
        val P = (1..N).map { input_ints() }

        // good problem: greedy on sub-problem
        fun set2(): Long{
            val C = hashMapOf<Pair<Int, Int>, Long>()
            val D = 500
            for((x,y,c) in P){
                C[y to x] = c.toLong()
            }
            val R = MutableList(D+1){ MutableList(D+1){0L} }
            val L = MutableList(D+1) { MutableList(D+1){0L} }
            for(y in D downTo 0){
                if(y == D){
                    for(x in 0 ..D){
                        R[y][x] = (if(x==0) 0 else R[y][x-1]) + (C[y to x] ?: 0L)
                    }
                    for(x in 0 ..D)
                        L[y][x] = R[0][D]-E
                }else {
                    for(x in 0 ..D){
                        // bug: if().. else .. + b --> (if()..else..) + b
                        R[y][x] = (if(x == 0){ maxOf(L[y+1][x]-E, R[y+1][x]) }else maxOf(R[y+1][x], R[y][x-1])) + (C[y to x] ?: 0L)
                    }
                    for(x in D downTo  0){
                        L[y][x] = (if(x == D){ maxOf(R[y+1][x]-E, L[y+1][x]) }else maxOf(L[y+1][x], L[y][x+1])) + (C[y to x] ?: 0L)
                    }
                }
            }
            var ma = (0..D).maxOf {y -> (0..D).maxOf {x -> maxOf(R[y][x], L[y][x]) } }
            return ma
        }

        println("Case #${it}: ${set2()}")
    }
}
