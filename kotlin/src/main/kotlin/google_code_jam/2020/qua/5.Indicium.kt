package google_code_jam.`2020`.qua

fun solve(N: Int, K: Int): String {
    val cube = Array(N){ Array(N){it+1} }
    val cols = Array(N){ Array(N+1){ false } }

    fun dfs(i: Int, j: Int): Boolean{
        if(i == N){
            val trace = (0 until N).sumOf { cube[it][it] }
            if(trace == K)
                return true;
        }else{
            val (ni, nj) = if(j==N-1) i+1 to 0 else i to j+1
            for(k in j until N) {
                val v = cube[i][k]
                if (!cols[j][v]) {
                    cols[j][v] = true
                    cube[i][j] = cube[i][k].also{cube[i][k] = cube[i][j]}
                    if(dfs(ni, nj))
                        return true;
                    cube[i][j] = cube[i][k].also{cube[i][k] = cube[i][j]}
                    cols[j][v] = false
                }
            }
        }
        return false
    }

    if(dfs(0, 0)){
        return buildString{
            appendLine("POSSIBLE")
            for(i in 0 until N)
                appendLine(cube[i].joinToString(" "))
        }
    }
    return "IMPOSSIBLE"
}

fun main(){
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, K) = readLine()!!.split(" ").map{it.toInt()}
        println("Case #${it}: ${solve(N, K)}")
    }
}