package google_code_jam.`2018`.`1b`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, L) = input_ints()
        val C = input_ints()
        var r = N - C.sum()
        var sum = 0
        val need = mutableListOf<Pair<Int, Int>>()
        fun cal(x: Int) = (x*200+N)/(2*N)
        for(c in C){
            sum += cal(c)
            val x = c*200 % (2*N)
            if(x >= N)
                continue  // bug: continue not break!!!
            for(t in 1..r){
                if((c+t)*200 % (2*N) >= N){
                    need.add(t to cal(c+t) - cal(c))
                    break
                }
            }
        }

        for(t in 1..r){
            if(t*200 % (2*N) >= N){
                for(k in 0 until r/t){ // bug: shouldn't decrease r here
                    need.add(t to cal(t))
                }
                break
            }
        }
        need.sortWith(compareBy({it.first}, {-it.second}))
//        println(C)
//        println(sum)
//        println(need)
        for((k, diff) in need){
            if(r < k)
                break
            r -= k
            sum += diff
        }
        sum += cal(r)
        println("Case #${it}: $sum")
    }
}
