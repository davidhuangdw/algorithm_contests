package google_kickstart.`2020`.F
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val line = input_ints()
        val (S, ai, aj, bi, bj) = line
        val blocked = List(S+1){BooleanArray(2*S)}
        blocked[ai][aj] = true
        blocked[bi][bj] = true
        repeat(line[5]){
            val (i, j) = input_ints()
            blocked[i][j] = true
        }

        val adj = List(S+1){ List(2*S){mutableListOf<Pair<Int, Int>>()}}
        fun valid(i: Int, j: Int) = i in 1..S && j in 1 until 2*i
        for(i in 1..S)
            for(j in 1 until 2*i){
                with(adj[i][j]){
                    if(j-1 in (1 until 2*i))
                        add(i to j-1)
                    if(j+1 in (1 until 2*i))
                        add(i to j+1)
                    if(j % 2 == 0 && valid(i-1, j-1)) add(i-1 to j-1)
                    if(j % 2 == 1 && valid(i+1, j+1)) add(i+1 to j+1)

                }
            }

        fun cal(xi: Int, xj: Int, yi: Int, yj: Int, done_x: Boolean, done_y: Boolean): Int{
            if(done_x && done_y) return 0
            if(done_x) return -cal(yi, yj, xi, xj, done_y, done_x)

            var res = Int.MIN_VALUE
            for((i, j) in adj[xi][xj]){
                if(!blocked[i][j]){
                    blocked[i][j] = true
                    res = maxOf(res, 1-cal(yi, yj, i, j, done_y, done_x))
                    blocked[i][j] = false
                }
            }
            if(res == Int.MIN_VALUE)
                res = -cal(yi, yj, xi, xj, done_y, true)
            return res
        }


        println("Case #${it}: ${cal(ai, aj, bi, bj, false, false)}")
    }
}
