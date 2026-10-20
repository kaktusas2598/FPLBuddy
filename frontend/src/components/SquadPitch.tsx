import type { Squad } from '../types.ts';

function SquadPitch({ squad }: { squad: Squad}) {

    return (
        <>
            {squad && (
                <div className="manager-result">
                    <h2>Squad</h2>
                    <ul>
                        {squad.picks.map(({ player, pick }) => (
                            <li key={player.id}>
                                {pick.position}. {player.web_name}
                                {pick.is_captain && ' (C)'}
                                {pick.is_vice_captain && ' (VC)'}
                            </li>
                        ))}
                    </ul>
                </div>
            )}

        </>
    )
}

export default SquadPitch

