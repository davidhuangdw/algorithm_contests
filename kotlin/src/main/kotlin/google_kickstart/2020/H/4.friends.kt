package google_kickstart.`2020`.H
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val MAX = 1e9.toInt()
    val M = 26
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, Q) = input_ints()
        val names = split_input()
        val dis = Array(M){ IntArray(M){MAX} }

        for(i in 0 until N) {
            for(a in names[i])
                for(b in names[i])
                    dis[a-'A'][b-'A'] = 1
        }

        for(i in 0 until M)
            dis[i][i] = 0

        for(k in 0 until M)
            for(i in 0 until M)
                for(j in 0 until M)
                    dis[i][j] = minOf(dis[i][j], dis[i][k]+dis[k][j])

        val ans = mutableListOf<Int>()
        repeat(Q){
            var (x, y) = input_ints()
            var min = MAX
            for(a in names[x-1])
                for(b in names[y-1]) {
                    min = minOf(min, dis[a - 'A'][b - 'A'])
                }
            ans.add(if(min == MAX) -1 else min+2)
        }

        println("Case #${it}: ${ans.joinToString(" ")}")
    }
}
